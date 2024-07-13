import { SavedFlow } from "@prisma/client";
import { parseRoute, parseRoutes } from ".";
import cache from "../lib/cache";
import db from "../lib/db";
import { savedFlowSchema } from "../lib/schemas";
import { TSavedFlow, TSavedFlow_createRequest, TSavedFlow_updateRequest } from "../lib/types";
import { initMakeRedisKey } from "../lib/utils";

const makeKey = initMakeRedisKey("flow");

export async function getAllFlows(): Promise<TSavedFlow[]> {
    const savedFlows: SavedFlow[] = await db.savedFlow.findMany();
    const proms: Promise<TSavedFlow>[] = savedFlows.map(flow => makeClientFlow(flow));
    return Promise.all(proms);
}

export async function getFlowById(id: number): Promise<TSavedFlow | null> {
    // Check redis cache for this flow
    const key = makeKey(id);
    const cachedResult = await cache?.get(key);

    // If found in the cache, parse and return it
    if (cachedResult != null) {
        const { data, success } = await savedFlowSchema.safeParseAsync(cachedResult);
        if (success) return data;
    }

    // If not in cache, query db for it
    const flowProm = db.savedFlow.findUnique({
        where: { id },
    });

    // If we fetch from the db successfully, create a new key for this flow in the cache
    flowProm.then(flow => {
        if (flow && cache) {
            cache.set(key, JSON.stringify(flow), {
                EX: 60, // Exipry time in seconds
            });
        }
    });

    return flowProm.then(flow => flow ? makeClientFlow(flow) : null);
}

export async function createNewFlow(creationRequest: TSavedFlow_createRequest): Promise<TSavedFlow> {
    const flowProm = db.savedFlow.create({
        data: {
            ...creationRequest,
            // Changing mainRoute and ruleRoutes into strings because
            // they are saved as strings in the db
            mainRoute: JSON.stringify(creationRequest.mainRoute),
            ruleRoutes: JSON.stringify(creationRequest.ruleRoutes),
        }
    });

    // If the creation was successful, create a new key for this new flow in the cache
    flowProm.then(flow => {
        if (flow && cache) {
            const key = makeKey(flow.id);
            cache.set(key, JSON.stringify(flow), {
                EX: 60, // Exipry time in seconds
            });
        }
    });

    return flowProm.then(flow => makeClientFlow(flow));
}

export async function updateFlowById(id: number, data: TSavedFlow_updateRequest): Promise<TSavedFlow> {
    const flowProm = db.savedFlow.update({
        where: { id },
        data: {
            ...data,
            // Changing mainRoute and ruleRoutes into strings because
            // they are saved as strings in the db
            mainRoute: JSON.stringify(data.mainRoute),
            ruleRoutes: JSON.stringify(data.ruleRoutes),
        }
    });

    // If the update was successful, update the corresponding key for this flow in the cache
    flowProm.then(flow => {
        if (flow && cache) {
            const key = makeKey(flow.id);
            cache.set(key, JSON.stringify(flow), {
                EX: 60, // Exipry time in seconds
            });
        }
    });

    return flowProm.then(flow => makeClientFlow(flow));
}

export async function deleteFlowById(id: number): Promise<TSavedFlow> {
    // Delete the corresponding key for this flow in the cache
    if (cache) {
        const key = makeKey(id);
        cache.del(key);
    }

    const flowProm = db.savedFlow.delete({
        where: { id },
    });

    return flowProm.then(flow => makeClientFlow(flow));
}

async function makeClientFlow(dbModel: SavedFlow): Promise<TSavedFlow> {
    const mainRouteProm = parseRoute(dbModel.mainRoute);
    const ruleRoutesProm = parseRoutes(dbModel.ruleRoutes);
    return {
        ...dbModel,
        mainRoute: await mainRouteProm,
        ruleRoutes: await ruleRoutesProm,
    };
}

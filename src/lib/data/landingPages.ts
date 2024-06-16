
import cache from '../cache';
import db from '../db';
import { landingPageSchema } from '../schemas';
import { LandingPage, LandingPage_createRequest, LandingPage_updateRequest } from '../types';
import { initMakeRedisKey } from '../utils';

const makeKey = initMakeRedisKey('landingPage');

export async function getAllLandingPages(): Promise<LandingPage[]> {
    return db.landingPage.findMany();
}

export async function getLandingPageById(id: number): Promise<LandingPage | null> {
    // Check redis cache for this landing page
    const key = makeKey(id);
    const cachedResult = await cache?.get(key);

    // If found in the cache, parse and return it
    if (cachedResult != null) {
        const { data, success } = await landingPageSchema.safeParseAsync(cachedResult);
        if (success) return data;
    }

    // If not in cache, query db for it
    const landingPageProm = db.landingPage.findUnique({
        where: { id }
    });

    // If we fetch from the db successfully, create a new key for this landing page in the cache
    landingPageProm.then(landingPage => {
        if (landingPage && cache) {
            cache.set(key, JSON.stringify(landingPage), {
                EX: 60 // Exipry time in seconds
            });
        }
    });

    return landingPageProm;
}

export async function createNewLandingPage(landingPageReqest: LandingPage_createRequest): Promise<LandingPage> {
    const landingPageProm = db.landingPage.create({
        data: { ...landingPageReqest }
    });

    // If the creation was successful, create a new key for this new landing page in the cache
    landingPageProm.then(landingPage => {
        if (landingPage && cache) {
            const key = makeKey(landingPage.id);
            cache.set(key, JSON.stringify(landingPage), {
                EX: 60 // Exipry time in seconds
            });
        }
    });

    return landingPageProm;
}

export async function updateLandingPageById(id: number, data: LandingPage_updateRequest): Promise<LandingPage> {
    const landingPageProm = db.landingPage.update({
        where: { id },
        data
    });

    // If the update was successful, update the corresponding key for this landing page in the cache
    landingPageProm.then(landingPage => {
        if (landingPage && cache) {
            const key = makeKey(landingPage.id);
            cache.set(key, JSON.stringify(landingPage), {
                EX: 60 // Exipry time in seconds
            });
        }
    });

    return landingPageProm;
}

export async function deleteLandingPageById(id: number): Promise<LandingPage> {
    // Delete the corresponding key for this landing page in the cache
    if (cache) {
        const key = makeKey(id);
        cache.del(key);
    }

    return db.landingPage.delete({
        where: { id }
    });
}

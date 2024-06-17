import { z } from 'zod';
import { toZod } from 'tozod';
import { TAffiliateNetwork, TLandingPage, TOffer, TToken, TTrafficSource, TUser } from './types';

export const userSchema: toZod<TUser> = z.object({
    id: z.number(),
    name: z.string(),
    hashedPassword: z.string(),
    createdAt: z.date(),
    updatedAt: z.date()
});

export const affiliateNetworkSchema: toZod<TAffiliateNetwork> = z.object({
    id: z.number(),
    name: z.string(),
    defaultNewOfferString: z.string(),
    tags: z.array(z.string()),
    createdAt: z.date(),
    updatedAt: z.date()
});

export const landingPageSchema: toZod<TLandingPage> = z.object({
    id: z.number(),
    name: z.string(),
    url: z.string(),
    tags: z.array(z.string()),
    createdAt: z.date(),
    updatedAt: z.date()
});

export const offersSchema: toZod<TOffer> = z.object({
    id: z.number(),
    name: z.string(),
    url: z.string(),
    payout: z.number(),
    tags: z.array(z.string()),
    createdAt: z.date(),
    updatedAt: z.date(),
    affiliateNetworkId: z.number()
});

export const tokenSchema: toZod<TToken> = z.object({
    queryParam: z.string(),
    value: z.string(),
    name: z.string()
});

export const trafficSourceSchema: toZod<TTrafficSource> = z.object({
    id: z.number(),
    name: z.string(),
    postbackUrl: z.string(),
    defaultTokens: z.array(tokenSchema),
    customTokens: z.array(tokenSchema),
    tags: z.array(z.string()),
    createdAt: z.date(),
    updatedAt: z.date()
});

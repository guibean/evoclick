"use client";

import {
    faBullseye, faFolder, faGlobe, faGlobeEurope, faHandshake,
    faLaptop, faMobile, faSitemap, faUsers, faWifi
} from "@fortawesome/free-solid-svg-icons";
import { EItemName } from "@/lib/types";
import UpperCPWrapper from "./UpperCPWrapper";
import UpperCPRow from "./UpperCPRow";
import UpperCPItem, { TUpperCPItem } from "./UpperCPItem";
import UpperCPItemGroup, { TUpperCPItemGroup } from "./UpperCPItemGroup";

export default function UpperControlPanel({ onClick = () => { }, reportItemName }: {
    onClick?: (item: TUpperCPItem) => void;
    reportItemName?: EItemName;
}) {
    return (
        <UpperCPWrapper>
            <UpperCPRow>
                {upperCPItems
                    .filter(item => item.itemName !== reportItemName)
                    .map((item, _index) => (
                        <UpperCPItem
                            key={_index}
                            item={item}
                            onClick={onClick}
                        />
                    ))}
            </UpperCPRow>
            <UpperCPRow>
                {upperCPItemGroups.map((group, index) => (
                    <UpperCPItemGroup
                        key={index}
                        itemGroup={group}
                        onClick={onClick}
                        reportItemName={reportItemName}
                    />
                ))}
            </UpperCPRow>
        </UpperCPWrapper>
    )
}

export const upperCPItems: TUpperCPItem[] = [
    { itemName: EItemName.CAMPAIGN, icon: faBullseye },
    { itemName: EItemName.OFFER, icon: faHandshake },
    { itemName: EItemName.LANDING_PAGE, icon: faFolder },
    { itemName: EItemName.FLOW, icon: faSitemap },
    { itemName: EItemName.TRAFFIC_SOURCE, icon: faGlobe },
    { itemName: EItemName.AFFILIATE_NETWORK, icon: faUsers },
];

export const upperCPItemGroups: TUpperCPItemGroup[] = [
    {
        name: EItemName.COUNTRY, icon: faGlobeEurope, children: [
            { itemName: EItemName.COUNTRY, icon: faGlobeEurope },
            { itemName: EItemName.REGION, icon: faGlobeEurope },
            { itemName: EItemName.CITY, icon: faGlobeEurope },
            { itemName: EItemName.LANGUAGE, icon: faGlobeEurope },
        ],
    },
    {
        name: EItemName.ISP, icon: faWifi, children: [
            { itemName: EItemName.ISP, icon: faWifi },
            { itemName: EItemName.IP, icon: faWifi },
            { itemName: EItemName.USER_AGENT, icon: faWifi },
        ],
    },
    {
        name: EItemName.DEVICE, icon: faLaptop, children: [
            { itemName: EItemName.DEVICE, icon: faLaptop },
            { itemName: EItemName.DEVICE_TYPE, icon: faLaptop },
            { itemName: EItemName.SCREEN_RESOLUTION, icon: faLaptop },
        ],
    },
    {
        name: EItemName.OS, icon: faMobile, children: [
            { itemName: EItemName.OS, icon: faMobile },
            { itemName: EItemName.OS_VERSION, icon: faMobile },
        ],
    },
    {
        name: EItemName.BROWSER_NAME, icon: faFolder, children: [
            { itemName: EItemName.BROWSER_NAME, icon: faFolder },
            { itemName: EItemName.BROWSER_VERSION, icon: faFolder },
        ],
    },
];

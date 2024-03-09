import {Pack} from "./pack";

export interface ShippingItem {
    pack: Pack
    count: number
}

export interface ShippingResponse {
    shipping: {
        items: ShippingItem[]
    }
    text: string
}
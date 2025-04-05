import { api } from "@/lib/axios";

export interface CreateColorRequest {
    storeId: string;
    name: string;
    hex: string;
}

export async function createColor({ storeId, name, hex }: CreateColorRequest) {
     await api.post(`/stores/${storeId}/colors`, {
        name,
        hex,
    });

}
import { api } from "@/lib/axios";

export interface UpdateStoreRequest {
    storeId: string;
    name: string;
}

export async function updateStore({ storeId, name }: UpdateStoreRequest) {
    await api.put(`/stores/${storeId}`, {
        name,
    });

}
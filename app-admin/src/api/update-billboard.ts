import { api } from "@/lib/axios";

export interface UpdateBillboardRequest {
    label: string;
    image: File;
    billboardId: string;
    storeId: string;
}

export async function updateBillboard({ label, image, billboardId, storeId }: UpdateBillboardRequest) {
    const formData = new FormData();
    formData.append("label", label);
    formData.append("image", image);

    await api.put(`/stores/${storeId}/billboards/${billboardId}`, formData);
}
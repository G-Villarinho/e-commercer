import { api } from "@/lib/axios";

export interface CreateBillboardRequest {
  label: string;
  image: File;
  storeId: string;
}

export async function createBillboard({ label, image, storeId }: CreateBillboardRequest) {
  const formData = new FormData();
  formData.append("label", label);
  formData.append("image", image); 

  await api.post(`/stores/${storeId}/billboards`, formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
}

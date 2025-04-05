import { api } from "@/lib/axios";
import { ColorResponse } from "./responses/colors-response";

export interface GetColorByIdParams {
  storeId: string;
  colorId: string;
}

export async function getColorById({ storeId, colorId }: GetColorByIdParams) {
  const response = await api.get<ColorResponse>(
    `/stores/${storeId}/colors/${colorId}`
  );

  return response.data;
}

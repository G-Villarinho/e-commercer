import { BillboardResponse } from "./billboard-response";

export interface CategoryResponse {
  id: string;
  name: string;
  createdAt: string;
  billboard: BillboardResponse;
}

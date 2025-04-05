import { api } from "@/lib/axios";

export interface CreateProductRequest {
  storeId: string;
  name: string;
  price: number;
  isArchived: boolean;
  isFeatured: boolean;
  categoryId: string;
  colorId: string;
  sizeId: string;

  images: File[];
}

export async function createProduct({
  storeId,
  name,
  price,
  isArchived,
  isFeatured,
  categoryId,
  colorId,
  sizeId,
  images,
}: CreateProductRequest) {
  const formData = new FormData();

  formData.append("name", name);
  formData.append("price", price.toString());
  formData.append("isArchived", isArchived.toString());
  formData.append("isFeatured", isFeatured.toString());
  formData.append("categoryId", categoryId);
  formData.append("colorId", colorId);
  formData.append("sizeId", sizeId);

  images.forEach((image) => {
    formData.append("images", image);
  });

  const response = await api.post(`/stores/${storeId}/products`, formData);

  return response.data;
}

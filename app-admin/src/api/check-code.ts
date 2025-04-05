import { api } from "@/lib/axios";

export async function checkCode() {
    return await api.get("/check-code");
}
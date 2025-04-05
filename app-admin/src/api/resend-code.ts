import { api } from "@/lib/axios";

export async function resendCode() {
    await api.post("/resend-code");
}
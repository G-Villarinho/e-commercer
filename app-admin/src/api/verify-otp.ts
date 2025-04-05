import { api } from "@/lib/axios";

export interface VerifyOtpRequest {
    code: string;
}

export async function VerifyOtp({ code }: VerifyOtpRequest) {
    return await api.post("/verify-code", { code });
}
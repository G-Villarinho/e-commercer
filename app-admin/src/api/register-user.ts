import { api } from "@/lib/axios";

export interface RegisterUserRequest {
    name: string;
    email: string;
}

export async function registerUser({name, email}: RegisterUserRequest) {
    await api.post("/register", {name, email});
}
import { AuthUser } from "../models";
import { ApiService } from "./api";

export class AuthService {
    private api: ApiService

    constructor(baseUrl: string, authToken?: string) {
        const headers: Record<string, string> = authToken ? { Authorization: `Bearer ${authToken}`} : {}
        this.api = new ApiService(baseUrl, headers)
    }

    async login(data: {user: string; password: string}): Promise<AuthUser> {
        return this.api.post('user/login', data)
    }
}
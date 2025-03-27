import { Pagination, Product } from "../models";
import { ApiService } from "./api";

export class ProductsService {
    private api: ApiService

    constructor(baseUrl: string, authToken?: string) {
        const headers: Record<string, string> = authToken ? { Authorization: `Bearer ${authToken}`} : {}
        this.api = new ApiService(baseUrl, headers)
    }

    async getProductsList(params: {page: number, pageSize: number, search?: string}): Promise<Pagination<Product>> {
        return this.api.get('product', params)
    }
}
import { Pagination, Product, ProductDTO, ProductStockDTO } from "../models";
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

    async getProductById(id: string): Promise<Product> {
        return this.api.get(`product/${id}`)
    }

    async createProduct(data: ProductDTO): Promise<Product> {
        return this.api.post('product', data)
    }

    async updateProduct(data: ProductDTO): Promise<Product> {
        return this.api.patch('product', data)
    }

    async updateStock(data: ProductStockDTO): Promise<Product> {
        return this.api.put('product/stock', data)
    }
}
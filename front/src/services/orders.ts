import { Pagination, Order, OrderDTO, OrderType, SummaryDash } from "../models";
import { ApiService } from "./api";

export interface OrderParams {page: number, pageSize: number, orderType: OrderType, search?: string}

export class OrdersService {
    private api: ApiService

    constructor(baseUrl: string, authToken?: string) {
        const headers: Record<string, string> = authToken ? { Authorization: `Bearer ${authToken}`} : {}
        this.api = new ApiService(baseUrl, headers)
    }

    async getOrdersList(params: {page: number, pageSize: number, search?: string}): Promise<Pagination<Order>> {
        return this.api.get('order', params)
    }

    async getOrderById(id: string): Promise<Order> {
        return this.api.get(`order/${id}`)
    }

    async createOrder(data: OrderDTO): Promise<Order> {
        return this.api.post('order', data)
    }

    async updateOrder(data: OrderDTO): Promise<Order> {
        return this.api.patch('order', data)
    }

    async getSummary(): Promise<SummaryDash> {
        return this.api.get('order/summary')
    }
}
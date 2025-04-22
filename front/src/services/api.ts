import { ApiError } from "../models";

type HttpMethod = 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE';

interface FetchOptions {
  method?: HttpMethod;
  headers?: Record<string, string>;
  body?: any;
  params?: Record<string, string | number | boolean>;
}

export class ApiService {
  private baseUrl: string;
  private defaultHeaders: Record<string, string>;

  constructor(baseUrl: string, defaultHeaders: Record<string, string> = {}) {
    this.baseUrl = baseUrl;
    this.defaultHeaders = {
      ...defaultHeaders,
    };
  }

  private async request<T>(endpoint: string, options: FetchOptions = {}): Promise<T> {
    const { method = 'GET', headers = {}, body, params } = options;

    const url = new URL(`${this.baseUrl}/${endpoint}`);
    if (params) {
      Object.keys(params).forEach(key => 
        url.searchParams.append(key, String(params[key]))
      );
    }

    const response = await fetch(url.toString(), {
      method,
      headers: {
        ...this.defaultHeaders,
        ...headers,
        "Content-Type": "application/json"
      },
      body: body ? JSON.stringify(body) : undefined,
    });

    if (!response.ok) {
      const errorData = await this.parseError(response);
      throw new ApiError({
        code: errorData.code,
        message: errorData.message,
        statusCode: response.status,
      });
    }

    return this.parseResponse<T>(response);
  }

  private async parseResponse<T>(response: Response): Promise<T> {
    const contentType = response.headers.get('Content-Type');
    if (contentType && contentType.includes('application/json')) {
      return response.json();
    }
    return response.text() as unknown as T;
  }

  private async parseError(response: Response): Promise<ApiError> {
    try {
      return await response.json();
    } catch {
      return { message: response.statusText, code: "ERR-GEN100", statusCode: 500 };
    }
  }

  // Métodos HTTP específicos para mejor semántica
  public get<T>(endpoint: string, params?: Record<string, string | number | boolean>, headers?: Record<string, string>): Promise<T> {
    return this.request<T>(endpoint, { method: 'GET', params, headers });
  }

  public post<T>(endpoint: string, body?: any, headers?: Record<string, string>): Promise<T> {
    return this.request<T>(endpoint, { method: 'POST', body, headers });
  }

  public put<T>(endpoint: string, body?: any, headers?: Record<string, string>): Promise<T> {
    return this.request<T>(endpoint, { method: 'PUT', body, headers });
  }

  public patch<T>(endpoint: string, body?: any, headers?: Record<string, string>): Promise<T> {
    return this.request<T>(endpoint, { method: 'PATCH', body, headers });
  }

  public delete<T>(endpoint: string, headers?: Record<string, string>): Promise<T> {
    return this.request<T>(endpoint, { method: 'DELETE', headers });
  }
}
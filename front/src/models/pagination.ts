export type Metadata = {
		page: number;
		pageSize: number;
		count: number;
		hasNext: boolean;
}

export type Pagination<T> = {
    metadata: Metadata
    data: T[]
}
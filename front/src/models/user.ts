export type Store = {
    id: string
	name: string
	contact:string
	address:string
	city:string
	status:string
	createdAt:string
	updatedAt:string
}

export type User = {
    id: string
    name: string
    email: string
    status: "ACTIVE" | "INACTIVE"
    role: "SUDO" | "ADMIN" | "MANAGER"
    store: Store
    createdAt:string
    updatedAt:string
}

export type AuthUser = User & {token: string}
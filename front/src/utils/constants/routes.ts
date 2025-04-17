export const ROUTES = {
    root: "/",

    products: "/products",
    ProductsAvailability: "/products/availability",
    productsCreate: "/products/create",

    orders: "/orders",
    ordersCreate: "/orders/create",

    login: "/login"
}

export const PRIVATE_ROUTES = [ROUTES.root, ROUTES.orders, ROUTES.products]
import { BrowserRouter, Route, Routes } from 'react-router'
import AuthLayout from './Layout/AuthLayout.tsx'
import LoginContainer from './modules/login/container/index.tsx'
import MainLayout from './Layout/MainLayout.tsx'
import OrdersList from './modules/orders/containers/OrdersList.tsx';
import OrderCreate from './modules/orders/containers/OrderCreate.tsx';
import OrderUpdate from './modules/orders/containers/OrderUpdate.tsx';
import ProductsList from './modules/products/containers/ProductsList.tsx';
import ProductCreate from './modules/products/containers/ProductCreate.tsx';
import ProductUpdate from './modules/products/containers/ProductUpdate.tsx';
import MainRoute from './modules/main.tsx';
import ProductsAvailability from './modules/products/containers/ProductsAvailability.tsx';

function App() {

    return (
        <BrowserRouter>
            <Routes>
                <Route element={<MainLayout />}>
                    <Route index element={<MainRoute />} />

                    <Route path="products" >
                        <Route index element={<ProductsList />} />
                        <Route path='create' element={<ProductCreate />} />
                        <Route path=':productId' element={<ProductUpdate />} />
                        <Route path='availability' element={<ProductsAvailability />} />
                    </Route>

                    <Route path="orders">
                        <Route index element={<OrdersList />} />
                        <Route path='create' element={<OrderCreate />} />
                        <Route path=':orderId' element={<OrderUpdate />} />
                    </Route>
                </Route>

                <Route element={<AuthLayout />}>
                    <Route path="login" element={<LoginContainer />} />
                </Route>
            </Routes>
        </BrowserRouter>
     );
}

export default App;
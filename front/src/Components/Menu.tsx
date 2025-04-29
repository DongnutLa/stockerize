import { Menu, MenuProps } from "antd";
import {
    AppstoreOutlined,
    MailOutlined,
    PieChartOutlined,
    HomeOutlined,
  } from '@ant-design/icons';
import { NavigateFunction, useNavigate } from "react-router";
import { ROUTES } from "../utils/constants";

type MenuItem = Required<MenuProps>['items'][number];

const items = (navigate: NavigateFunction): MenuItem[] => [
    { key: 'home', icon: <HomeOutlined />, label: 'Inicio', onClick: () => navigate(ROUTES.root) },
    {
      key: 'products',
      label: 'Productos',
      icon: <MailOutlined />,
      children: [
        { key: 'list-products', label: 'Lista de productos', onClick: () => navigate(ROUTES.products), },
        { key: 'create-products', label: 'Crear productos', onClick: () => navigate(ROUTES.productsCreate), },
        { key: 'products-availability', label: 'Disponibilidad', onClick: () => navigate(ROUTES.ProductsAvailability), },
      ],
    },
    {
      key: 'orders',
      label: 'Órdenes',
      icon: <AppstoreOutlined />,
      children: [
        { key: 'list-orders', label: 'Lista de órdenes', onClick: () => navigate(ROUTES.orders), },
        { key: 'create-order', label: 'Crear orden', onClick: () => navigate(ROUTES.ordersCreate), },
      ],
    },
    { key: 'logout', icon: <PieChartOutlined />, label: 'Cerrar sesión' },
  ];

function MainMenu({collapsed}: {collapsed: boolean }) {
  let navigate = useNavigate();

    return (
        <Menu
          theme={"dark"}
          defaultSelectedKeys={['1']}
          defaultOpenKeys={['sub1']}
          mode="inline"
          inlineCollapsed={collapsed}
          items={items(navigate)}
        />
    )
}

export default MainMenu
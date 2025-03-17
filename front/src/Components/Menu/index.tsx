import { Menu, MenuProps } from "antd";
import {
    AppstoreOutlined,
    ContainerOutlined,
    DesktopOutlined,
    MailOutlined,
    PieChartOutlined,
  } from '@ant-design/icons';

type MenuItem = Required<MenuProps>['items'][number];

const items: MenuItem[] = [
    {
      key: 'products',
      label: 'Productos',
      icon: <MailOutlined />,
      children: [
        { key: 'list-products', label: 'Lista de productos' },
        { key: 'create-products', label: 'Crear productos' },
      ],
    },
    {
      key: 'orders',
      label: 'Órdenes',
      icon: <AppstoreOutlined />,
      children: [
        { key: 'list-orders', label: 'Lista de órdenes' },
        { key: 'create-order', label: 'Crear orden' },
      ],
    },
    { key: 'logout', icon: <PieChartOutlined />, label: 'Cerrar sesión' },
  ];

function MainMenu() {
    return <Menu
    defaultSelectedKeys={['1']}
    defaultOpenKeys={['sub1']}
    mode="inline"
    items={items}
  />
}

export default MainMenu
//import FC from react
import type { FC } from "react";

//import Link from react router dom
import { Link } from "react-router";

//import custom hook useLogout
import { useLogout } from "../hooks/auth/useLogout";

const SidebarMenu: FC = () => {

    //initialize useLogout
    const logout = useLogout();

    return (
        <aside className="sidebar-card">
            <div className="sidebar-heading">
                Workspace
            </div>
            <nav className="sidebar-nav" aria-label="Admin navigation">
                    <Link to="/admin/dashboard" className="list-group-item list-group-item-action">Dashboard</Link>

                    <Link to="/admin/users" className="list-group-item list-group-item-action">Users</Link>
                    <Link to="/admin/products" className="list-group-item list-group-item-action">Products</Link>
                    <Link to="/admin/categories" className="list-group-item list-group-item-action">Categories</Link>
                    <Link to="/admin/suppliers" className="list-group-item list-group-item-action">Suppliers</Link>

                    <a onClick={logout} className="list-group-item list-group-item-action" style={{ cursor: 'pointer' }}>Logout</a>
            </nav>
        </aside>
    )
}

export default SidebarMenu;

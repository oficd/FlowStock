//import useContext
import { useContext } from 'react';

//import context
import { AuthContext } from '../context/AuthContext.tsx';

//import react router dom
import { Routes, Route, Navigate } from "react-router";

//import view home
import Home from "../views/home/index.tsx";

//import view register
import Register from "../views/auth/register.tsx";

//import view login
import Login from "../views/auth/login.tsx";

//import view dashboard
import Dashboard from "../views/admin/dashboard/index.tsx";

//import view users
import UsersIndex from "../views/admin/users/index.tsx";

//import view users create
import UsersCreate from "../views/admin/users/create.tsx";

//import view users edit
import UsersEdit from "../views/admin/users/edit.tsx";

//import view products
import ProductsIndex from "../views/admin/products/index.tsx";

//import view products create
import ProductsCreate from "../views/admin/products/create.tsx";

//import view products edit
import ProductsEdit from "../views/admin/products/edit.tsx";

//import categories
import CategoriesIndex from "../views/admin/categories/index.tsx";
import CategoriesCreate from "../views/admin/categories/create.tsx";

//import suppliers
import SuppliersIndex from "../views/admin/suppliers/index.tsx";
import SuppliersCreate from "../views/admin/suppliers/create.tsx";

export default function AppRoutes() {

    // Menggunakan useContext untuk mendapatkan nilai dari AuthContext
    const auth = useContext(AuthContext);

    // Menggunakan optional chaining untuk menghindari error jika auth tidak ada
    const isAuthenticated = auth?.isAuthenticated ?? false;

    return (
        <Routes>
            {/* route "/" */}
            <Route path="/" element={<Home />} />

            {/* route "/register" */}
            <Route path="/register" element={
                isAuthenticated ? <Navigate to="/admin/dashboard" replace /> : <Register />
            } />

            {/* route "/login" */}
            <Route path="/login" element={
                isAuthenticated ? <Navigate to="/admin/dashboard" replace /> : <Login />
            } />

            {/* route "/admin/dashboard" */}
            <Route path="/admin/dashboard" element={
                isAuthenticated ? <Dashboard /> : <Navigate to="/login" replace />
            } />

            {/* route "/admin/users" */}
            <Route path="/admin/users" element={
                isAuthenticated ? <UsersIndex /> : <Navigate to="/login" replace />
            } />

            {/* route "/admin/users/create" */}
            <Route path="/admin/users/create" element={
                isAuthenticated ? <UsersCreate /> : <Navigate to="/login" replace />
            } />

            {/* route "/admin/users/edit/:id" */}
            <Route path="/admin/users/edit/:id" element={
                isAuthenticated ? <UsersEdit /> : <Navigate to="/login" replace />
            } />

            {/* route "/admin/products" */}
            <Route path="/admin/products" element={
                isAuthenticated ? <ProductsIndex /> : <Navigate to="/login" replace />
            } />

            {/* route "/admin/products/create" */}
            <Route path="/admin/products/create" element={
                isAuthenticated ? <ProductsCreate /> : <Navigate to="/login" replace />
            } />

            {/* route "/admin/products/edit/:id" */}
            <Route path="/admin/products/edit/:id" element={
                isAuthenticated ? <ProductsEdit /> : <Navigate to="/login" replace />
            } />

            {/* route "/admin/categories" */}
            <Route path="/admin/categories" element={
                isAuthenticated ? <CategoriesIndex /> : <Navigate to="/login" replace />
            } />

            {/* route "/admin/categories/create" */}
            <Route path="/admin/categories/create" element={
                isAuthenticated ? <CategoriesCreate /> : <Navigate to="/login" replace />
            } />

            {/* route "/admin/suppliers" */}
            <Route path="/admin/suppliers" element={
                isAuthenticated ? <SuppliersIndex /> : <Navigate to="/login" replace />
            } />

            {/* route "/admin/suppliers/create" */}
            <Route path="/admin/suppliers/create" element={
                isAuthenticated ? <SuppliersCreate /> : <Navigate to="/login" replace />
            } />
        </Routes>
    );
}

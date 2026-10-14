//import FC from react
import type { FC } from "react";

//import SidebarMenu
import SidebarMenu from '../../../components/SidebarMenu';

//import Link from react-route
import { Link } from "react-router";

//import custom hook useProducts
import { useProducts } from "../../../hooks/product/useProducts";
import type { Product } from "../../../hooks/product/useProducts";

//import custom hook useProductDelete
import { useProductDelete } from "../../../hooks/product/useProductDelete";

//import query client TanStack Query
import { useQueryClient } from '@tanstack/react-query';

const ProductsIndex: FC = () => {

    // get products from useProducts
    const { data: products, isLoading, isError, error } = useProducts();

    //initialize useQueryClient
    const queryClient = useQueryClient();

    //initialize useProductDelete
    const { mutate, isPending } = useProductDelete();

    //handle delete product
    const handleDelete = (id: number) => {
        if(confirm("Are you sure you want to delete this product?")) {
            
            //call useProductDelete
            mutate(id, {
                onSuccess: () => {
                    //refetch data
                    queryClient.invalidateQueries({ queryKey: ['products'] });
                }
            });
        }
    }

    return (
        <div className="container mt-5 mb-5">
            <div className="row">
                <div className="col-md-3">
                    <SidebarMenu />
                </div>
                <div className="col-md-9">
                    <div className="card border-0 rounded-4 shadow-sm">
                        <div className="card-header d-flex justify-content-between align-items-center">
                            <span>PRODUCTS</span>
                            <Link to="/admin/products/create" className="btn btn-sm btn-success rounded-4 shadow-sm border-0">ADD PRODUCT</Link>
                        </div>
                        <div className="card-body">

                            {/* Loading State */}
                            {isLoading && (
                                <div className="alert alert-info text-center">Loading...</div>
                            )}

                            {/* Error State */}
                            {isError && (
                                <div className="alert alert-danger text-center">
                                    Error: {error.message}
                                </div>
                            )}

                            <table className="table table-bordered">
                                <thead className="bg-dark text-white">
                                    <tr>
                                        <th scope="col">Product Name</th>
                                        <th scope="col">Category</th>
                                        <th scope="col">Supplier</th>
                                        <th scope="col">Total</th>
                                        <th scope="col" style={{ width: "20%" }}>Actions</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {
                                        products?.map((product: Product) => (
                                            <tr key={product.id}>
                                                <td>{product.name}<div className="text-muted small">{product.description}</div></td>
                                                <td>{product.category?.name ?? (product.category_id ?? '-')}</td>
                                                <td>{product.supplier?.name ?? (product.supplier_id ?? '-')}</td>
                                                <td>{product.total ?? 0}</td>
                                                <td className="text-center">
                                                    <Link to={`/admin/products/edit/${product.id}`} className="btn btn-sm btn-primary rounded-4 shadow-sm border-0 me-2">EDIT</Link>
                                                    <button onClick={() => handleDelete(product.id)} disabled={isPending} className="btn btn-sm btn-danger rounded-4 shadow-sm border-0">
                                                        {isPending ? 'DELETING...' : 'DELETE'}
                                                    </button>
                                                </td>
                                            </tr>
                                        ))
                                    }
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default ProductsIndex

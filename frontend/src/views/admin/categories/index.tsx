import type { FC } from "react";
import SidebarMenu from '../../../components/SidebarMenu';
import { Link } from "react-router";
import { useCategories } from '../../../hooks/category/useCategories';
import type { Category } from '../../../hooks/category/useCategories';
import { useCategoryDelete } from '../../../hooks/category/useCategoryDelete';
import { useQueryClient } from '@tanstack/react-query';

const CategoriesIndex: FC = () => {

    const { data: categories, isLoading, isError, error } = useCategories();
    const queryClient = useQueryClient();
    const { mutate, isPending } = useCategoryDelete();

    const handleDelete = (id: number) => {
        if(confirm("Are you sure you want to delete this category?")) {
            mutate(id, {
                onSuccess: () => {
                    queryClient.invalidateQueries({ queryKey: ['categories'] });
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
                            <span>CATEGORIES</span>
                            <Link to="/admin/categories/create" className="btn btn-sm btn-success rounded-4 shadow-sm border-0">ADD CATEGORY</Link>
                        </div>
                        <div className="card-body">

                            {isLoading && (
                                <div className="alert alert-info text-center">Loading...</div>
                            )}

                            {isError && (
                                <div className="alert alert-danger text-center">Error: {error.message}</div>
                            )}

                            <table className="table table-bordered">
                                <thead className="bg-dark text-white">
                                    <tr>
                                        <th scope="col">Category Name</th>
                                        <th scope="col" style={{ width: "20%" }}>Actions</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {categories?.map((c: Category) => (
                                        <tr key={c.id}>
                                            <td>{c.name}</td>
                                            <td className="text-center">
                                                <Link to={`/admin/categories/edit/${c.id}`} className="btn btn-sm btn-primary rounded-4 shadow-sm border-0 me-2">EDIT</Link>
                                                <button onClick={() => handleDelete(c.id)} disabled={isPending} className="btn btn-sm btn-danger rounded-4 shadow-sm border-0">{isPending ? 'DELETING...' : 'DELETE'}</button>
                                            </td>
                                        </tr>
                                    ))}
                                </tbody>
                            </table>

                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default CategoriesIndex;

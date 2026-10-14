import type { FC } from "react";
import SidebarMenu from '../../../components/SidebarMenu';
import { Link } from "react-router";
import { useSuppliers } from '../../../hooks/supplier/useSuppliers';
import type { Supplier } from '../../../hooks/supplier/useSuppliers';
import { useSupplierDelete } from '../../../hooks/supplier/useSupplierDelete';
import { useQueryClient } from '@tanstack/react-query';

const SuppliersIndex: FC = () => {

    const { data: suppliers, isLoading, isError, error } = useSuppliers();
    const queryClient = useQueryClient();
    const { mutate, isPending } = useSupplierDelete();

    const handleDelete = (id: number) => {
        if(confirm("Are you sure you want to delete this supplier?")) {
            mutate(id, {
                onSuccess: () => {
                    queryClient.invalidateQueries({ queryKey: ['suppliers'] });
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
                            <span>SUPPLIERS</span>
                            <Link to="/admin/suppliers/create" className="btn btn-sm btn-success rounded-4 shadow-sm border-0">ADD SUPPLIER</Link>
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
                                        <th scope="col">Supplier Name</th>
                                        <th scope="col" style={{ width: "20%" }}>Actions</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {suppliers?.map((s: Supplier) => (
                                        <tr key={s.id}>
                                            <td>{s.name}</td>
                                            <td className="text-center">
                                                <Link to={`/admin/suppliers/edit/${s.id}`} className="btn btn-sm btn-primary rounded-4 shadow-sm border-0 me-2">EDIT</Link>
                                                <button onClick={() => handleDelete(s.id)} disabled={isPending} className="btn btn-sm btn-danger rounded-4 shadow-sm border-0">{isPending ? 'DELETING...' : 'DELETE'}</button>
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

export default SuppliersIndex;

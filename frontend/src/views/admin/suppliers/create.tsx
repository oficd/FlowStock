import { useState } from "react";
import type { FC, FormEvent } from "react";
import SidebarMenu from '../../../components/SidebarMenu';
import { useNavigate, Link } from 'react-router';
import { useSupplierCreate } from '../../../hooks/supplier/useSupplierCreate';

interface ValidationErrors {
    [key: string]: string;
}

const SupplierCreate: FC = () => {

    const navigate = useNavigate();
    const [name, setName] = useState<string>('');
    const [errors, setErrors] = useState<ValidationErrors>({});
    const { mutate, isPending } = useSupplierCreate();

    const storeSupplier = async (e: FormEvent) => {
        e.preventDefault();

        mutate({ name }, {
            onSuccess: () => {
                navigate('/admin/suppliers');
            },
            onError: (error: any) => {
                setErrors(error.response.data.errors);
            }
        })
    }

    return (
        <div className="container mt-5 mb-5">
            <div className="row">
                <div className="col-md-3">
                    <SidebarMenu />
                </div>
                <div className="col-md-9">
                    <div className="card border-0 rounded-4 shadow-sm">
                        <div className="card-header">ADD SUPPLIER</div>
                        <div className="card-body">
                            <form onSubmit={storeSupplier}>

                                <div className="form-group mb-3">
                                    <label className="mb-1 fw-bold">Supplier Name</label>
                                    <input type="text" value={name} onChange={(e) => setName(e.target.value)} className="form-control" placeholder="Supplier Name" />
                                    {errors.Name && <div className="alert alert-danger mt-2 rounded-4">{errors.Name}</div>}
                                </div>

                                <button type="submit" className="btn btn-md btn-primary rounded-4 shadow-sm border-0" disabled={isPending}>{isPending ? 'Saving...' : 'Save'}</button>
                                <Link to="/admin/suppliers" className="btn btn-md btn-secondary rounded-4 shadow-sm border-0 ms-2">Cancel</Link>

                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default SupplierCreate;

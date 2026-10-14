//import FC from react
import { useState } from "react";
import type { FC, FormEvent } from "react";

//import SidebarMenu
import SidebarMenu from '../../../components/SidebarMenu';

//import useNavigate and Link from react-router
import { useNavigate, Link } from 'react-router';

//import custom hook useProductCreate
import { useProductCreate } from '../../../hooks/product/useProductCreate';
import { useCategories } from '../../../hooks/category/useCategories';
import { useSuppliers } from '../../../hooks/supplier/useSuppliers';

//interface for validation errors
interface ValidationErrors {
    [key: string]: string;
}

const ProductCreate: FC = () => {

    //initialize useNavigate
    const navigate = useNavigate();

    //define state product
    const [name, setName] = useState<string>('');
    const [description, setDescription] = useState<string>('');
    const [total, setTotal] = useState<number>(0);
    const [categoryId, setCategoryId] = useState<number | ''>('');
    const [supplierId, setSupplierId] = useState<number | ''>('');

    const { data: categories } = useCategories();
    const { data: suppliers } = useSuppliers();

    //define state errors
    const [errors, setErrors] = useState<ValidationErrors>({});

    // Inisialisasi useProductCreate
    const { mutate, isPending } = useProductCreate();

    // Handle submit form
    const storeProduct = async (e: FormEvent) => {
        e.preventDefault();

        // Call the product create mutation
        mutate({
            name,
            description,
            total,
            category_id: categoryId === '' ? undefined : Number(categoryId),
            supplier_id: supplierId === '' ? undefined : Number(supplierId),
        }, {
            onSuccess: () => {

                // Redirect to products index
                navigate('/admin/products');
            },
            onError: (error: any) => {

                //set errors to state "errors"
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
                        <div className="card-header">
                            ADD PRODUCT
                        </div>
                        <div className="card-body">
                            <form onSubmit={storeProduct}>


                                <div className="form-group mb-3">
                                    <label className="mb-1 fw-bold">Product Name</label>
                                    <input type="text" value={name} onChange={(e) => setName(e.target.value)} className="form-control" placeholder="Product Name" />
                                    {errors.Name && <div className="alert alert-danger mt-2 rounded-4">{errors.Name}</div>}
                                </div>

                                <div className="form-group mb-3">
                                    <label className="mb-1 fw-bold">Description</label>
                                    <textarea value={description} onChange={(e) => setDescription(e.target.value)} className="form-control" placeholder="Description" />
                                    {errors.Description && <div className="alert alert-danger mt-2 rounded-4">{errors.Description}</div>}
                                </div>

                                <div className="form-group mb-3">
                                    <label className="mb-1 fw-bold">Total</label>
                                    <input type="number" value={total} onChange={(e) => setTotal(Number(e.target.value))} className="form-control" placeholder="Total" />
                                    {errors.Total && <div className="alert alert-danger mt-2 rounded-4">{errors.Total}</div>}
                                </div>

                                <div className="form-group mb-3">
                                    <label className="mb-1 fw-bold">Category</label>
                                    <select value={categoryId} onChange={(e) => setCategoryId(e.target.value === '' ? '' : Number(e.target.value))} className="form-control">
                                        <option value="">-- Select Category --</option>
                                        {categories?.map((cat) => (
                                            <option key={cat.id} value={cat.id}>{cat.name}</option>
                                        ))}
                                    </select>
                                    {errors.CategoryId && <div className="alert alert-danger mt-2 rounded-4">{errors.CategoryId}</div>}
                                </div>

                                <div className="form-group mb-3">
                                    <label className="mb-1 fw-bold">Supplier</label>
                                    <select value={supplierId} onChange={(e) => setSupplierId(e.target.value === '' ? '' : Number(e.target.value))} className="form-control">
                                        <option value="">-- Select Supplier --</option>
                                        {suppliers?.map((s) => (
                                            <option key={s.id} value={s.id}>{s.name}</option>
                                        ))}
                                    </select>
                                    {errors.SupplierId && <div className="alert alert-danger mt-2 rounded-4">{errors.SupplierId}</div>}
                                </div>

                                <button type="submit" className="btn btn-md btn-primary rounded-4 shadow-sm border-0" disabled={isPending}>
                                    {isPending ? 'Saving...' : 'Save'}
                                </button>

                                <Link to="/admin/products" className="btn btn-md btn-secondary rounded-4 shadow-sm border-0 ms-2">Cancel</Link>

                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div >
    )
}

export default ProductCreate

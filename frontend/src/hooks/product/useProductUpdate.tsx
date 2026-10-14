// import useMutation dari '@tanstack/react-query';
import { useMutation } from '@tanstack/react-query';

//import service API
import Api from '../../services/api';

// import js-cookie
import Cookies from 'js-cookie';

//interface ProductRequest untuk update
interface ProductRequest {
    name: string;
    description?: string;
    total?: number;
    category_id?: number;
    supplier_id?: number;
}

export const useProductUpdate = () => {
    return useMutation({
        mutationFn: async ({ id, data }: { id: number, data: ProductRequest }) => {
            const token = Cookies.get('token');

            const response = await Api.put(`/api/products/${id}`, data, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data;
        }
    });
};

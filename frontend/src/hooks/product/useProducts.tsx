// import hook useQuery from react-query
import { useQuery } from '@tanstack/react-query';

//import service Api
import Api from '../../services/api';

// import js-cookie
import Cookies from 'js-cookie';

//interface Product
export interface Product {
    id: number;
    name: string;
    description?: string;
    total?: number;
    category_id?: number;
    category?: { id: number; name: string };
    supplier_id?: number;
    supplier?: { id: number; name: string };
    created_at?: string;
    updated_at?: string;
}

//hook useProducts dengan return type Product
export const useProducts = () => {

    return useQuery<Product[], Error>({

        //query key
        queryKey: ['products'],

        //query function
        queryFn: async () => {

            //get token from cookies
            const token = Cookies.get('token');

            //get products from api
            const response = await Api.get('/api/products', {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            
            //return data
            return response.data.data as Product[];
        },
        
    });
}

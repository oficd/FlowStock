// import hook useQuery from react-query
import { useQuery } from '@tanstack/react-query';

//import service Api
import Api from '../../services/api';

// import js-cookie
import Cookies from 'js-cookie';

//interface Supplier
export interface Supplier {
    id: number;
    name: string;
}

//hook useSuppliers
export const useSuppliers = () => {

    return useQuery<Supplier[], Error>({

        queryKey: ['suppliers'],

        queryFn: async () => {

            const token = Cookies.get('token');

            const response = await Api.get('/api/suppliers', {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data.data as Supplier[];
        }

    });
}

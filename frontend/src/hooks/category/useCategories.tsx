// import hook useQuery from react-query
import { useQuery } from '@tanstack/react-query';

//import service Api
import Api from '../../services/api';

// import js-cookie
import Cookies from 'js-cookie';

//interface Category
export interface Category {
    id: number;
    name: string;
}

//hook useCategories
export const useCategories = () => {

    return useQuery<Category[], Error>({

        queryKey: ['categories'],

        queryFn: async () => {

            const token = Cookies.get('token');

            const response = await Api.get('/api/categories', {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data.data as Category[];
        }

    });
}

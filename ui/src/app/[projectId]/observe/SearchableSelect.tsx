// import { useEffect, useState } from 'react';
// import { Select, SelectItem } from '@nextui-org/react';
// import {fetchLangChainFilterValues} from "@/services/langchain_runs";
//
//
// interface SearchableSelectProps {
//     projectId: string
//     labelKey: string
//     condition: string
//     index: number
//     handleChange: (index: number, field: string, value: string) => void;
// }
//
// function SearchableSelect({projectId, labelKey, condition, index, handleChange}: SearchableSelectProps) {
//     const [options, setOptions] = useState([]);
//     const [loading, setLoading] = useState(false);
//     const [page, setPage] = useState(1);
//
//     const fetchOptions = async (searchTerm: string, page: number) => {
//         setLoading(true);
//         try {
//             const data = await fetchLangChainFilterValues(projectId, labelKey, condition, searchTerm, page, 20);
//             setOptions((prevOptions: any[]) => [...prevOptions, ...data]);
//             setLoading(false);
//         } catch (error) {
//             console.error('Error fetching options:', error);
//             setLoading(false);
//         }
//     };
//
//     const handleSearch = (value: any) => {
//         setOptions([]);  // Clear the options
//         setPage(1);      // Reset to the first page
//         if (value) {
//             fetchOptions(value, 1);
//         }
//     };
//
//     const handleInfiniteScroll = () => {
//         if (loading) return; // Avoid triggering multiple times if already loading
//         setPage(prevPage => prevPage + 1);
//     };
//
//     useEffect(() => {
//         if (page > 1) {
//             // Assuming the Select component provides the current search term
//             const selector = document.querySelector('.nextui-select-input')
//             if (selector !== null) {
//                 const currentSearchTerm = selector.value;
//                 fetchOptions(currentSearchTerm, page).then(value => {
//
//                 });
//             }
//
//         }
//     }, [page]);
//
//     return (
//         <Select
//             color={"primary"}
//             placeholder="Search values ..."
//             onChange={(e) => handleChange(index, 'value', e.target.value)}
//             // onSearch={handleSearch} TODO: Figure out handling search
//             // onInfinite={handleInfiniteScroll} TODO: Figure out handling search
//             loading={loading}
//             items={options}
//         >
//             {(option: any) => (
//                 <SelectItem key={option.id} value={option.id}>
//                     {option.name}
//                 </SelectItem>
//             )}
//         </Select>
//     );
// }
//
// export default SearchableSelect;

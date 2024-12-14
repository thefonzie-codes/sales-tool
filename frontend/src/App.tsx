import './App.css'
import { useState } from 'react'

async function fetchCustomer(customer: string) {
    const api = import.meta.env.VITE_BACKEND_API // Ensure you're using the correct environment variable
    console.log(import.meta.env.VITE_BACKEND_API); 

    try {
        const res = await fetch(`${api}/customers?name=${customer}`);
        console.log(`Fetching: ${api}/customers?name=${customer}`);
        if (!res.ok) {
            throw new Error(`Error: ${res.status} - ${res.statusText}`)
        }

        const data = await res.json();
        return data;
    } catch (err) {
        console.error("Error fetching customer data: ", err)
        throw err;
    }
}

function App() {
    interface FormData {
        name: string;
    }

    const [formData, setFormData] = useState<FormData>({
        name: ""
    });

    const [customerData, setCustomerData] = useState<any>(null); // State to store fetched customer data
    const [error, setError] = useState<string | null>(null); // State to store any error messages
    const [loading, setLoading] = useState(false); // State to track loading status

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
        console.log(formData);
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setLoading(true);  // Start loading state

        try {
            const data = await fetchCustomer(formData.name);  // Fetch customer data
            setCustomerData(data);  // Update state with fetched data
            setError(null);  // Clear any previous error
        } catch (err) {
            setError('Error fetching customer data');
            setCustomerData(null);  // Clear previous data if an error occurs
        } finally {
            setLoading(false);  // End loading state
        }
    };

    return (
        <>
            <div>
                <p>Search for customer</p>
                <form onSubmit={handleSubmit}>
                    <input
                        type="text"
                        name="name"
                        value={formData.name}
                        onChange={handleChange}
                        placeholder="Enter customer name"
                    />
                    <button type="submit" disabled={loading}>
                        {loading ? 'Loading...' : 'Submit'}
                    </button>
                </form>

                {error && <p style={{ color: 'red' }}>{error}</p>} {/* Display error message */}

                {customerData && (
                    <div>
                        <h3>Customer Details</h3>
                        <p><strong>Name:</strong> {customerData.first_name} {customerData.last_name}</p>
                        <p><strong>Account Name:</strong> {customerData.account_name}</p>
                        <p><strong>Company:</strong> {customerData.company}</p>
                        <p><strong>Email:</strong> {customerData.email}</p>
                        <p><strong>Phone:</strong> {customerData.phone}</p>
                        <p><strong>Address:</strong> {customerData.address}</p>
                    </div>
                )}
            </div>
        </>
    );
}

export default App;


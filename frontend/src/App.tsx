// import FC from react
import type { FC } from 'react';

//import router
import AppRoutes from './routes';

//import Link from react router
import { Link } from "react-router";

//import custom hook useLogout
// import { useLogout } from "./hooks/auth/useLogout";

const App: FC = () => {

  // //initialize useLogout
  // const logout = useLogout();

  return (
    <div>
      <nav className="navbar navbar-expand-lg app-navbar">
        <div className="container">
           <Link to="/" className="navbar-brand">Flow<span>Stock</span></Link>
          <button
            className="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarSupportedContent"
            aria-controls="navbarSupportedContent"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarSupportedContent">
            {/* <ul className="navbar-nav ms-auto mb-2 mb-lg-0">
               <li className="nav-item">
                 <a onClick={logout} className="nav-link active" style={{ cursor: 'pointer' }}>Logout</a>
              </li>
            </ul> */}
          </div>
        </div>
      </nav>

      <main className="app-main">
        <AppRoutes />
      </main>
    </div>
  )
}

export default App;

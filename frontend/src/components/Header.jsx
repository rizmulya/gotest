import { Link } from "react-router-dom";

const Header = () => {
  return (
    <div className="mb-5">
      Header
      <Link to="/" className="btn btn-primary">Home</Link>
      <Link to="/login" className="btn btn-primary">Login</Link>
      <Link to="/register" className="btn btn-primary">Register</Link>
      <Link to="/users" className="btn btn-primary">Users</Link>
    </div>
  );
};

export default Header;
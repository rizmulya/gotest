import { useEffect, useState, useRef } from 'react';
import api from '@/api';

const Index = () => {
  const [users, setUsers] = useState([]);
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [role, setRole] = useState('');
  const [image, setImage] = useState(null);
  const imageInputRef = useRef(null);

  const [error, setError] = useState('');

  const [editingUser, setEditingUser] = useState(null);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await api.get('/api/users');
        setUsers(response.data);
      } catch (error) {
        console.error(error);
        setError('Failed to fetch users');
      }
    };
    fetchUsers();
  }, []);

  const handleCreate = async (e) => {
    e.preventDefault();
    try {
      const formData = new FormData();
      formData.append('name', name);
      formData.append('email', email);
      formData.append('password', password);
      formData.append('role', role);
      if (image) {
        formData.append('image', image);
      }

      const response = await api.post('/api/users', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });

      if (response.status === 200) {
        setUsers([...users, response.data]);
        setName('');
        setEmail('');
        setPassword('');
        setRole('');
        setImage(null);
        imageInputRef.current.value = '';
      }
    } catch (error) {
      setError(error.response?.data?.error || 'Failed to create user');
    }
  };

  const handleUpdate = async (e) => {
    e.preventDefault();
    try {
      const formData = new FormData();
      formData.append('name', name);
      formData.append('email', email);
      if (password) {
        formData.append('password', password);
      }
      formData.append('role', role);
      if (image) {
        formData.append('image', image);
      }

      const response = await api.put(`/api/users/${editingUser.uid}`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });

      if (response.status === 200) {
        setUsers(users.map(user => (user.uid === editingUser.uid ? response.data : user)));
        setEditingUser(null);
        setName('');
        setEmail('');
        setPassword('');
        setRole('');
        setImage(null);
        imageInputRef.current.value = '';
      }
    } catch (error) {
      setError(error.response?.data?.error || 'Failed to update user');
    }
  };

  const handleDelete = async (uid) => {
    try {
      await api.delete(`/api/users/${uid}`);
      setUsers(users.filter(user => user.uid !== uid));
    } catch (error) {
      setError(error.response?.data?.error || 'Failed to delete user');
    }
  };

  const startEditUser = (user) => {
    setEditingUser(user);
    setName(user.name);
    setEmail(user.email);
    setPassword('');
    setRole(user.role);
    setImage(null);
    imageInputRef.current.value = '';
  };

  return (
    <div>
      <h2>Users</h2>
      <form onSubmit={editingUser ? handleUpdate : handleCreate}>
        <div>
          <label>Name</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Email</label>
          <input
            type="text"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Password</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required={!editingUser}
          />
        </div>
        <div>
          <label>Role</label>
          <input
            type="text"
            value={role}
            onChange={(e) => setRole(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Image</label>
          <input
            type="file"
            onChange={(e) => setImage(e.target.files[0])}
            ref={imageInputRef}
          />
        </div>
        {error && <div style={{ color: 'red' }}>{error}</div>}
        <button type="submit">{editingUser ? 'Update' : 'Create'}</button>
      </form>
      <br />
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Email</th>
            <th>Role</th>
            <th>Image</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {users.map(user => (
            <tr key={user.ID}>
              <td>{user.ID}</td>
              <td>{user.name}</td>
              <td>{user.email}</td>
              <td>{user.role}</td>
              <td>
                {user.image && (
                  <img src={`/static/uploads/images/${user.image}`} alt="User" width="50" />
                )}
              </td>
              <td>
                <button onClick={() => startEditUser(user)}>Edit</button>
                <button onClick={() => handleDelete(user.uid)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Index;

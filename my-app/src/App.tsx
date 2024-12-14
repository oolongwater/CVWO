// src/App.tsx
import React, { useEffect, useState } from 'react';
import { Container, Typography, List, ListItem } from '@mui/material';

interface Data {
  id: number;
  name: string;
}

const App: React.FC = () => {
  const [data, setData] = useState<Data[]>([]);

  useEffect(() => {
    fetch('/api/data')
      .then(response => response.json())
      .then(response => setData(response.payload.data)); // Adjusted to handle nested structure
  }, []);

  return (
    <Container>
      <Typography variant="h4">Data from Backend</Typography>
      <List>
        {data.map(item => (
          <ListItem key={item.id}>{item.name}</ListItem>
        ))}
      </List>
    </Container>
  );
};

export default App;
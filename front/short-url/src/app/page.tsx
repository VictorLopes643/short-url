'use client';
import { useEffect, useState } from 'react';
import axios from 'axios';

import styles from './page.module.css';
import { useRouter } from 'next/navigation';

interface hashProps {
  params: {
    slug:string
  }
}

export default function Home({params}: hashProps) {

  const [inputValue, setInputValue] = useState('');
  const router = useRouter();
  const [res, setRes] = useState([]);

  const handleInputChange = (event:any) => {
    setInputValue(event.target.value);
  };

  // const handleSubmit = (event) => {
  //   event.preventDefault();
  //   console.log('Valor digitado:', inputValue);
  //   setInputValue('');
  // };

  const handleSubmit = async (event:any) => {
    event.preventDefault();

    try {
      const payload = {
        url: inputValue
      }
      const response = await axios.post('https://hujvceope9.execute-api.us-east-1.amazonaws.com/prod/urlHash', payload);

      console.log('Resposta da API:', response.data);

      // Faça algo com a resposta da API, se necessário

      // Limpe o valor do input
      setInputValue('');
      handleHash()
    } catch (error) {
      console.error('Erro ao chamar a API:', error);
    }
  };

  const handleHash = async () => {
    try {  
      // Solicitação OPTIONS para a API
      const optionsResponse = await axios.options('https://hujvceope9.execute-api.us-east-1.amazonaws.com/prod/operationsUrl');
      console.log("Resposta da solicitação OPTIONS:", optionsResponse);
      const response = await axios.get('https://hujvceope9.execute-api.us-east-1.amazonaws.com/prod/operationsUrl');
      console.log('Resposta da API:', response.data);
      setRes(response.data)
      setInputValue('');

    } catch (error) {
      console.error('Erro ao chamar a API:', error);
    }
  }

  const removeHash = async (id:string) => {
    try {  
      // Solicitação OPTIONS para a API
      const optionsResponse = await axios.options('https://hujvceope9.execute-api.us-east-1.amazonaws.com/prod/operationsUrl');
      console.log("Resposta da solicitação OPTIONS:", optionsResponse);
    
      const payload = {
        data: { id: id } 
      };
      
      const response = await axios.delete('https://hujvceope9.execute-api.us-east-1.amazonaws.com/prod/operationsUrl', payload);
      handleHash()
      console.log('Resposta da API:', response.data);
      setRes(response.data)
      setInputValue('');
    } catch (error) {
      console.error('Erro ao chamar a API:', error);
    }
  }


  useEffect(() => {
    handleHash()
  },[])
console.log("res:", res)
  return (
    <main className={styles.main}>
    <div className={styles.banner}>
      <p>Mercado Livre</p>
    </div>
    <div className={styles.content}>
      <form onSubmit={handleSubmit} className={styles.containerButton}>
        <input
          className={styles.search}
          value={inputValue}
          onChange={handleInputChange}
          placeholder="Digite um Link..."
        />
        <button type="submit" className={styles.button}>
          Enviar
        </button>
      </form>
      <div className={styles.tableContainer}>
        <table className={styles.mytable}>
          <thead>
            <tr>
              <th>Default URL</th>
              <th>Hash</th>
              <th>ID</th>
              <th>Remover</th>
            </tr>
          </thead>
          <tbody>
          {Array.isArray(res) && res.length > 0 ? (
            res.map((item: any, index: any) => (
              <tr key={index}>
                <td>{item.defaultUrl.S}</td>
                <td>{item.hash.S}</td>
                <td>{item.id.S}</td>
                <td>
                  <button className={styles.buttonRemove} onClick={() => removeHash(item.id.S)}>X</button>
                </td>
              </tr>
            ))
          ) : (
                <p>Adicione uma URL</p>
          )}
          </tbody>
        </table>
      </div>
    </div>
  </main>
  );
}
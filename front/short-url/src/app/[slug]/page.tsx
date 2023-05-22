'use client'
import { useEffect, useState } from 'react';
import axios from 'axios';

import styles from '../page.module.css';
import { redirect } from 'next/navigation';

interface hashProps {
  params: {
    slug:string
  }
}

export default function Home({params}: hashProps) {

  const [inputValue, setInputValue] = useState('');


  const handleInputChange = (event:any) => {
    setInputValue(event.target.value);
  };

  const handleSubmit = async (event:any) => {
    event.preventDefault();

    try {
      const payload = {
        url: inputValue
      }
      console.log("Payload:", payload)

      const response = await axios.post('https://hujvceope9.execute-api.us-east-1.amazonaws.com/prod/urlHash', payload);

      console.log('Resposta da API:', response.data);

  
      // Faça algo com a resposta da API, se necessário

      // Limpe o valor do input
      setInputValue('');
     
    } catch (error) {
      console.error('Erro ao chamar a API:', error);
    }
  };
  const handleHash = async () => {
    try {
      const payload = {
        hash: params.slug
      };
  
      console.log("Payload:", payload);
  
      // Solicitação OPTIONS para a API
      const optionsResponse = await axios.options('https://hujvceope9.execute-api.us-east-1.amazonaws.com/prod/redirectionUrl');
      console.log("Resposta da solicitação OPTIONS:", optionsResponse);
  
      // Verifique se a resposta OPTIONS inclui o cabeçalho 'Access-Control-Allow-Origin'
      // if (optionsResponse.headers['access-control-allow-origin']) {
        // Se o cabeçalho estiver presente, faça a solicitação POST
        const response = await axios.post('https://hujvceope9.execute-api.us-east-1.amazonaws.com/prod/redirectionUrl', payload);
        console.log('Resposta da API:', response.data);
        
        // Faça algo com a resposta da API, se necessário
  
        // Limpe o valor do input
        setInputValue('');
        const redirectTimeout = setTimeout(() => {
          window.location.href = response.data; // Substitua com o domínio de destino
        }, 3000);
    
        // Limpar o timeout quando o componente é desmontado
        return () => clearTimeout(redirectTimeout);
      // } else {
      //   console.error('A resposta OPTIONS não inclui o cabeçalho Access-Control-Allow-Origin');
      // }
    } catch (error) {
      console.error('Erro ao chamar a API:', error);
    }
  }

  useEffect(() => {
    handleHash()
  },[])

  return (
    <main className={styles.main}>
      <div className={styles.banner}>
        {/* <img
          src='../../public/mllogo.jpg'
          width={500}
          height={500}
          color=''
          alt="Picture of the author"
        /> */}
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
      </div>
    </main>
  );
}
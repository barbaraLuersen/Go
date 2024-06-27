 <h1>Projeto Go: Aplicação Web com PostgreSQL</h1>

<p>Este projeto é uma aplicação web desenvolvida como parte do curso de Go da plataforma Alura. A aplicação web criada utiliza a linguagem Go e integra um banco de dados PostgreSQL para realizar operações básicas de CRUD (Create, Read, Update, Delete), utilizando o framework Bootstrap para o frontend.</p>

<h2>Funcionalidades</h2>
   <ul>
        <li><strong>CRUD Completo</strong>: Capacidade de criar, ler, atualizar e deletar registros no banco de dados.</li>
        <li><strong>Interface Web</strong>: Utilização de templates HTML com estilos fornecidos pelo Bootstrap para renderizar páginas web dinâmicas.</li>
   </ul>

<h2>Tecnologias Utilizadas</h2>
        <ul>
            <li><strong>Go</strong>: Linguagem de programação principal.</li>
            <li><strong>PostgreSQL</strong>: Banco de dados relacional para armazenamento dos dados.</li>
            <li><strong>HTML/CSS</strong>: Frontend utilizando templates HTML e estilos CSS fornecidos pelo Bootstrap.</li>
            <li><strong>Bootstrap</strong>: Framework de frontend para criação de interfaces web responsivas.</li>
        </ul>

  <h2>Pré-requisitos</h2>
        <p>Antes de iniciar, certifique-se de ter instalado em sua máquina:</p>
        <ul>
            <li><a href="https://go.dev/doc/install">Go</a></li>
            <li><a href="https://www.postgresql.org/download/">PostgreSQL</a></li>
            <li><a href="https://getbootstrap.com/docs/4.3/getting-started/download/">Bootstrap</a></li>
        </ul>

<h2>Instalação e Uso</h2>
        <ol>
            <li>Clone este repositório: <code>git clone https://github.com/seu-usuario/nome-do-repositorio.git</code></li>
            <li>Acesse o diretório do projeto: <code>cd nome-do-repositorio</code></li>
            <li>Instale as dependências necessárias: <code>go mod tidy</code></li>
            <li>Configure o banco de dados PostgreSQL:</li>
            <ul>
                <li>Crie um banco de dados e as tabelas necessárias.</li>
                <li>Configure as variáveis de ambiente para conexão com o banco de dados (opcional).</li>
            </ul>
            <li>Compile e execute a aplicação: <code>go run main.go</code> ou <code>go build &amp;&amp; ./nome-do-executavel</code></li>
        </ol>

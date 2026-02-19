<a id="header"></a>

<center>
<a href="#header">
    <img src="./docs/assets/images/layout/header.png" alt="gopher azuis torcedores" />
</a>
</center>

<!-- 
    icons by:
    https://devicon.dev/
    https://simpleicons.org/
-->

[<img src="./docs/assets/images/icons/go.svg" width="25px" height="25px" alt="Go Logo" title="Go">](https://go.dev/) [<img src="./docs/assets/images/icons/postgresql.svg" width="25px" height="25px" alt="PostgreSql Logo" title="PostgreSql">](https://www.postgresql.org/) [<img src="./docs/assets/images/icons/docker.svg" width="25px" height="25px" alt="Docker Logo" title="Docker">](https://www.docker.com/) [<img src="./docs/assets/images/icons/ubuntu.svg" width="25px" height="25px Logo" title="Ubuntu" alt="Ubuntu" />](https://ubuntu.com/) [<img src="./docs/assets/images/icons/dotenv.svg" width="25px" height="25px" alt="Viper DotEnv Logo" title="Viper DotEnv">](https://github.com/spf13/viper) [<img src="./docs/assets/images/icons/github.svg" width="25px" height="25px" alt="GitHub Logo" title="GitHub">](https://github.com/jtonynet) [<img src="./docs/assets/images/icons/visualstudiocode.svg" width="25px" height="25px" alt="VsCode Logo" title="VsCode">](https://code.visualstudio.com/) [<img src="./docs/assets/images/icons/cursor.svg" width="25px" height="25px" alt="Cursor Logo" title="VsCode">](https://cursor.com/agents) 

[![Go Version](https://img.shields.io/badge/GO-1.24.11-blue?logo=go&logoColor=white)](https://go.dev/)

## 🕸️ Redes

[![linkedin](https://img.shields.io/badge/Linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/jos%C3%A9-r-99896a39/) [![gmail](https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:learningenuity@gmail.com)

---

## 📁 O Projeto

<a id="index"></a>
### ⤴️ Índice


__[Go Concurrent Importer](#header)__<br/>
  1.  ⤴️ [Índice](#index)
  2.  📖 [Sobre](#about)
  3.  💻 [Instalando](#install)
  4.  💻 [Rodando](#run)
  5.  ✅ [Testando](#tests)
  6.  🤖 [Uso de IA](#ia)

---

<a id="about"></a>
## 📖 Sobre

>  **Cenário:**
>  - Ler um arquivo CSV com 4 colunas: user_id, segment_type, segment_name e data
>  - Este arquivo terá 1 milhão de linhas
>  - O processamento deve ser performático e otimizado
>  - Validar se os dados são válidos
>  - Salvar no banco de dados
>  - Se houver erro, mostrar ou salvar essa informação de alguma forma
>
> ---

A aplicação Dockerizada foi testada em Sistema Operacional `Ubuntu 22.04.4 LTS`

<br/>

[⤴️ de volta ao índice](#index)

<br/>

<a id="install"></a>
## 💻 Instalando

`Docker` e `Docker Compose` são necessários para rodar a aplicação de forma containerizada, e é fortemente recomendado utilizá-los para rodar o banco de dados e demais dependências localmente. Siga as instruções abaixo caso não tenha esses softwares instalados em sua máquina:

- &nbsp;<img src='./docs/assets/images/icons/docker.svg' width='13' alt='Github do' title='Github do'>&nbsp;[Instalando Docker](https://docs.docker.com/engine/install/)
- &nbsp;<img src='./docs/assets/images/icons/docker.svg' width='13' alt='Github do' title='Github do'>&nbsp;[Instalando Docker Compose](https://docs.docker.com/compose/install/)


```bash
docker compose up -r
cd app
go mod tidy
```

<br/>

<a id="run"></a>
## 💻 Rodando
Com o docker rodando e a app instlada, digite:
```bash
cd app
go run ./cmd/cli/main.go
```

<br/>

<a id="tests"></a>
## ✅ Testando
```bash
cd app/internal/service/
go test -v -race
```

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="ia"></a>
### 🤖 Uso de IA

A figura do cabeçalho nesta página foi criada com a ajuda de inteligência artificial e um mínimo de retoques e construção no Gimp [<img src="./docs/assets/images/icons/gimp.svg" width="30" height="30 " title="Gimp" alt="Gimp Logo" />](https://www.gimp.org/)

__Os seguintes prompts foram usados para criação no  [Bing IA:](https://www.bing.com/images/create/)__

<details>
  <summary><b>Gopher Concorrendo em Maratona</b></summary>
"gophers azul, simbolo da linguagem golang com concorrendo em uma maratona, estilo cartoon, historia em quadrinhos, fundo branco chapado para facilitar remoção<b>(sic)</b>
</details>

<br/>

IA também é utilizada em minhas pesquisas e estudos como ferramenta de apoio em conjunto com a IDE Cursor [<img src="./docs/assets/images/icons/cursor.svg" width="25px" height="25px" alt="Cursor Logo" title="VsCode">](https://cursor.com/agents)

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="footer"></a>

<br/>

>  _"Lifelong Learning & Prosper"_
> <br/> 
>  _Mr. Spock, maybe_   🖖🏾🚀

<div align="center">
<a href="#footer">
<img src="./docs/assets/images/layout/footer.png" />
</a>
</div>

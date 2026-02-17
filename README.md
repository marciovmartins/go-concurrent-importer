# go-concurrent-importer

### About
Heavy concurrent data importer


>  **Cenário:**
>  - Ler um arquivo CSV com 4 colunas: user_id, segment_type, segment_name e data
>  - Este arquivo terá 1 milhão de linhas
>  - O processamento deve ser performático e otimizado
>  - Validar se os dados são válidos
>  - Salvar no banco de dados
>  - Se houver erro, mostrar ou salvar essa informação de alguma forma
>
> ---


<br/>



### Install
```bash
cd app
go mod tidy
```

<br/>

### Run
```bash
cd app
go run ./cmd/cli/main.go
```

<br/>

### Test
```bash
cd app/internal/service/
go test -v -race
```
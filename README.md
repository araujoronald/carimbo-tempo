
## Arquivos envolvidos

- acfts-bundle.pem (cadeia de certificados da ACT)
- go-data.txt (arquivo de texto com o conteúdo a ser assinado)
- tsr_go (arquivo de saída do resultado do carimbo)

## O que faz

Aplica um carimbo do tempo (timestamp) no conteúdo de um arquivo de texto e salva o resultado em um arquivo.

## Qual resultado

Após execução será gerado o arquivo tsr_go (resultado do carimbo).

## Como verificar o carimbo

```
openssl ts -verify -in tsr_go -data go-data.txt -CAfile acfts-bundle.pem
```

## Não foi possível recuperar um recurso
**Detalhe:** Houve um erro interno e não foi possível recuperar um recurso<br>
**Explicação:** Houve a tentativa de consultar um recurso no banco de dados, porém um erro desconhecido foi retornado.<br>
**Como resolver:** Será necessário criar uma issue para o time técnico realizar uma análise. No momento da criação da issue, recomendamos explicar detalhadamente o que foi realizado e enviar o conteúdo enviado na requisição, para facilitar o entendimento para os técnicos.

## Não foi possível cadastrar um recurso
**Detalhe:** Houve um erro interno e não foi possível cadastrar um recurso<br>
**Explicação:** Houve a tentativa de cadastrar um recurso no banco de dados, porém um erro desconhecido foi retornado.<br>
**Como resolver:** Será necessário criar uma issue para o time técnico realizar uma análise. No momento da criação da issue, recomendamos explicar detalhadamente o que foi realizado e enviar o conteúdo enviado na requisição, para facilitar o entendimento para os técnicos.

## Houve uma tentativa de reutilizar um dado único
**Detalhe:** O código informado já foi utilizado anteriormente<br>
**Explicação:** O atributo `code` é um atributo único dentro da aplicação utilizado para encontrar uma entidade. Esse atributo é considerado como um indentificador único dentro do sistema.<br>
**Como resolver:** 
 - Recomendamos que seja realizado uma analise no método utilizado para realizar a chamado do endpoint, pode ser que houve a tentativa de realizar uma atualização, mas por falta de atenção foi colocado o método POST no lugar do método PUT.
 - Se o registro existente com esse `code` não estiver correto, você pode realizar a atualização do mesmo para fazer uma regularização dos dados.
 - Cadastre esse recurso com um `code` diferente.

## Estrutura com dados inválidos
**Detalhe:** Estrutura com dados inválidos<br>
**Explicação:** A entidade que está sendo cadastrada não possui todos os dados corretos<br>
**Como resolver:** Analisar os modelos de criação da entidade em questão e comparar se as regras especificadas lá estão corretamente aplicadas no modelo que está tentando ser cadastrado.

## Não foi possível realizar a validação da estrutura
**Detalhe:** Não foi possível realizar a validação dos dados a serem cadastrados<br>
**Explicação:** Houve um erro não previsto quando realizada a tentativa de validar os dados da estrutura<br>
**Como resolver:** Será necessário criar uma issue para o time técnico realizar uma análise. No momento da criação da issue, recomendamos explicar detalhadamente o que foi realizado e enviar o conteúdo enviado na requisição, para facilitar o entendimento para os técnicos.

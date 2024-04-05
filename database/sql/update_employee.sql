update employees
set surname         = :surname,
    name            = :name,
    patronymic      = :patronymic,
    document_number = :document_number,
    birthday        = :birthday
where id = :id;

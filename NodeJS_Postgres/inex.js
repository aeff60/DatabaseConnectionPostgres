const express = require('express');
const { Client } = require('pg');
const connection = new Client({
          user: 'postgres',
          host: 'localhost',
          database: 'CompanyDB',
          password: 'mflv[DB2023',
          port: 5433,
});
connection.connect();         
const app = express();
app.get('/', (req, res) => {
          connection.query('SELECT * FROM public.employee', function (error, results, fields) {
                    if (error) throw error;
                    res.send(results.rows);
          }
          );
}
);
app.listen(3000, () => console.log('Listening on port 3000'));

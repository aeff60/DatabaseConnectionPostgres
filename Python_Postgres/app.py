import psycopg2
from flask import Flask, render_template, request, redirect, url_for, session, flash, jsonify
app = Flask(__name__)
conn = psycopg2.connect(
          host="localhost",
          port="5433",
          database="CompanyDB",
          user="postgres",
          password="mflv[DB2023"
)
cur = conn.cursor()
@app.route('/')
def index():
          cur.execute("SELECT * FROM public.employee")
          records = cur.fetchall()
          return jsonify(records)
app.run(debug=True, port=5000)
cur.close()
conn.close()
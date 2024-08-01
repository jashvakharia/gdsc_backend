from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///data.db'
db = SQLAlchemy(app)

class Data(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(50), nullable=False)
    value = db.Column(db.String(50), nullable=False)

with app.app_context():
    db.create_all()

@app.route('/test', methods=['GET'])
def get_sample():
    sample_response = {"message": "Welcome to the backend of GDSC."}
    return jsonify(sample_response)

@app.route('/post', methods=['POST'])
def post_data():
    data = request.get_json()
    name = data.get('name')
    value = data.get('value')

    if name and value:
        new_data = Data(name=name, value=value)
        db.session.add(new_data)
        db.session.commit()
        return jsonify({"message": "Your data has been saved successfully!"}), 201
    else:
        return jsonify({"message": "Invalid format/content-type."}), 400

@app.route('/data', methods=['GET'])
def get_data():
    data = Data.query.all()
    result = [
        {
            "id": d.id,
            "name": d.name,
            "value": d.value
        } for d in data
    ]
    return jsonify(result)

if __name__ == '__main__':
    app.run(debug=True)

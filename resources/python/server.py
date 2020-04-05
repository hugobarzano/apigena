from flask import render_template
import connexion
#https://realpython.com/flask-connexion-rest-api/
#https://github.com/go-swagger/go-swagger/blob/3596b40bedd4f06cc918767098568456d2ec2370/examples/2.0/petstore/server/api/petstore.go
# Create the application instance
app = connexion.App(__name__, specification_dir='../spec/')

# Read the swagger.yml file to configure the endpoints
app.add_api('../spec/specPi.yml')

@app.route('/')
def index():
    """
    This function just responds to the browser ULR
    :return:        the rendered template 'index.html'
    """
    return render_template('index.html')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3002, debug=True)
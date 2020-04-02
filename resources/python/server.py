from flask import render_template
import connexion
#https://realpython.com/flask-connexion-rest-api/
#https://github.com/go-swagger/go-swagger/blob/3596b40bedd4f06cc918767098568456d2ec2370/examples/2.0/petstore/server/api/petstore.go
# Create the application instance
app = connexion.App(__name__, specification_dir='../spec/')

# Read the swagger.yml file to configure the endpoints
app.add_api('../spec/specPi.yml')


# Create a URL route in our application for "/"
@app.route('/')
def home():
    """
    This function just responds to the browser ULR
    localhost:5000/
    :return:        the rendered template 'home.html'
    """
    return render_template('index.html')

# If we're running in stand alone mode, run the application
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3002, debug=True)
import socket
from flask import Flask

config = {
   "DEBUG": True
}
app = Flask(__name__)
app.config.from_mapping(config)

username="asdf"
pwd="zxcv@asdf"

@app.route('/')
def Home():
   hostname=socket.gethostname()
   ipAddr=socket.gethostbyname(hostname)
   return "<h1>Hello Jew777, "+hostname+", "+ipAddr+"</h1>"

@app.route('/mobileapp')
def mobileapp():
   return "<h1>This is return from mobile app.</h1>"

##### This part is for DAST testing #####
#@app.route('/9efc421d-d5e3-4a09-b1fe-6d3b63883378.html')
#def uuidchecking():
#   return ""

#@app.route('/forti-uuid.html')
#def uuidDetailCheck():
#   return "<forti-uuid hidden>9efc421d-d5e3-4a09-b1fe-6d3b63883378</forti-uuid>"

#####

if __name__ == '__main__':
 app.debug = True
 app.run(host='0.0.0.0', port=8000)

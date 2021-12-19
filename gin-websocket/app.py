# -*- coding: utf-8 -*-
# 자율주행 stack으로 부터 gps를 받는 부분을 구현해야 합니다.
# 이 코드는 고정된 gps를 사용합니다.
# 서버로 부터 DISPATCHED(배차됨), DRIVING(주행시작) 메시지를 받으면 route를 얻을 수 있습니다.
# 자율주행 stack에 목적지를 설정하는 부분을 구현해야 합니다.

import sys
import json
from collections import namedtuple
import websocket
from PyQt5.QtCore import *
from PyQt5.QtWidgets import *


GPS = namedtuple('GPS', 'lat lng speed bearing accuracy')


class WebsocketThread(QThread):

    on_message = pyqtSignal(str)

    def __init__(self, autoriaztion_key='1606367945:b57a77704caa500724520351988d6940'):
        #super().__init__()
        super(QThread, self).__init__()
        self.autoriaztion_key = autoriaztion_key

    def run(self):
        def on_open_handler(ws):
            print('on_open')

        def on_close_handler(ws):
            print('on_close')

        def on_message_handler(ws, message):
            self.on_message.emit(message)

        def on_error_handler(ws, err):
            print('on_error', err)

        #websocket.enableTrace(True)
        self.ws = websocket.WebSocketApp("ws://localhost:8080/ws",
                                    header={"Authorization": self.autoriaztion_key},
                                    on_open=on_open_handler,
                                    on_message=on_message_handler,
                                    on_error=on_error_handler,
                                    on_close=on_close_handler)
        self.ws.run_forever()

    @pyqtSlot(str)
    def send(self, msg):
        self.ws.send(msg)


class MyUI(QWidget):
    send_signal = pyqtSignal(str)

    def __init__(self):
        #super().__init__()
        super(MyUI, self).__init__()
        self.current_call = None
        self.current_route = None
        self.last_gps = GPS(37.393833, 127.112714, 10, 0, 10)
        self.layout_ui()

        # Signals:
        self.buttonPickup.clicked.connect(self.send_pickup)
        self.buttonNoshow.clicked.connect(self.send_noshow)
        self.buttonComplete.clicked.connect(self.send_dropoff)
        self.buttonTerminate.clicked.connect(self.send_terminate)

        self.timer = QTimer()
        self.timer.timeout.connect(self.send_gps)
        self.timer.start(2000)
        print('init')

    def layout_ui(self):
        self.setWindowTitle("자율주행 기사앱")
        self.setFixedWidth(500)
        self.setFixedHeight(300)

        layout = QVBoxLayout()
        layout.setContentsMargins(50, 50, 50, 50)
        layout.setSpacing(20)

        self.statusLabel = QLabel()
        self.statusLabel.setStyleSheet("font-weight: bold;")
        self.originLabel = QLabel()
        self.destLabel = QLabel()

        layout.addWidget(self.statusLabel)
        layout.addWidget(self.originLabel)
        layout.addWidget(self.destLabel)

        self.waitingFrame = QFrame()
        self.waitingFrame.setFixedHeight(100)

        self.pickupFrame = QFrame()
        self.pickupFrame.setFixedHeight(100)
        pickupButtons = QHBoxLayout()
        self.buttonPickup = QPushButton('승객탑승')
        self.buttonPickup.setStyleSheet("font-weight: bold;")
        self.buttonNoshow = QPushButton('noshow')
        self.buttonNoshow.setStyleSheet("color: #666;")
        pickupButtons.addWidget(self.buttonPickup)
        pickupButtons.addWidget(self.buttonNoshow)
        self.pickupFrame.setLayout(pickupButtons)

        self.drivingFrame = QFrame()
        self.drivingFrame.setFixedHeight(100)
        drivingButtons = QHBoxLayout()
        self.buttonComplete = QPushButton('운행완료')
        self.buttonComplete.setStyleSheet("font-weight: bold;")
        self.buttonTerminate = QPushButton('승객신고')
        self.buttonTerminate.setStyleSheet("color: #666;")
        drivingButtons.addWidget(self.buttonComplete)
        drivingButtons.addWidget(self.buttonTerminate)
        self.drivingFrame.setLayout(drivingButtons)

        layout.addWidget(self.waitingFrame)
        layout.addWidget(self.pickupFrame)
        layout.addWidget(self.drivingFrame)
        self.setLayout(layout)
        self.show()
        self.update_ui()

    def update_ui(self):
        if self.current_call is None:
            self.statusLabel.setText('콜대기')
            self.originLabel.setText('')
            self.destLabel.setText('')

            self.waitingFrame.show()
            self.pickupFrame.hide()
            self.drivingFrame.hide()
        else:
            if sys.version_info[0] < 3:
                self.originLabel.setText('출발지: ' + self.current_call['origin_name'].encode("utf-8"))
                self.destLabel.setText('도착지: ' + self.current_call['destination_name'].encode("utf-8"))
            else:
                self.originLabel.setText('출발지: ' + self.current_call['origin_name'])
                self.destLabel.setText('도착지: ' + self.current_call['destination_name'])

            if self.current_call['status'] == 'passenger_waiting':
                self.statusLabel.setText('픽업중')
                self.waitingFrame.hide()
                self.pickupFrame.show()
                self.drivingFrame.hide()
            elif self.current_call['status'] == 'driving':
                self.statusLabel.setText('주행중')
                self.waitingFrame.hide()
                self.pickupFrame.hide()
                self.drivingFrame.show()

    @pyqtSlot(str)
    def process_message(self, body):
        msg = json.loads(body)
        msg_type = msg['type']
        if msg_type in ('DISPATCHED', 'DRIVING'):
            self.current_call = msg['call']
            self.current_route = msg['route']
        else:
            self.current_call = None
        self.update_ui()

    def send(self, rec):
        msg = json.dumps(rec)
        print(msg)
        self.send_signal.emit(msg)

    def send_gps(self):
        rec = {'type': 'gps', 'lat': self.last_gps.lat, 'lng': self.last_gps.lng, 'accuracy': self.last_gps.accuracy,
               'speed': self.last_gps.speed, 'bearing': self.last_gps.bearing}
        self.send(rec)

    def send_pickup(self):
        self.send({'type': 'pickup', 'lat': self.last_gps.lat, 'lng': self.last_gps.lng, 'id': self.current_call['id']})

    def send_dropoff(self):
        self.send({'type': 'dropoff', 'lat': self.last_gps.lat, 'lng': self.last_gps.lng, 'id': self.current_call['id']})

    def send_noshow(self):
        self.send({'type': 'noshow', 'id': self.current_call['id']})

    def send_terminate(self):
        self.send({'type': 'terminate', 'id': self.current_call['id']})


def main(authorization_key):
    ws_thread = WebsocketThread(authorization_key)

    app = QApplication([])
    ui = MyUI()
    ws_thread.on_message.connect(ui.process_message)
    ui.send_signal.connect(ws_thread.send)
    ws_thread.start()

    app.exec_()


if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("python driverapp.py ${authorizationKey}")
        exit(1)
    main(sys.argv[1])

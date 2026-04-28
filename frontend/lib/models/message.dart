class Message {
  final String id;
  final String senderId;
  final String content;
  final int type;
  final DateTime timestamp;

  Message({
    required this.id,
    required this.senderId,
    required this.content,
    required this.type,
    required this.timestamp,
  });
}
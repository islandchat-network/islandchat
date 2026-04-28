import 'package:flutter/material.dart';
import 'screens/chat_screen.dart';

void main() {
  runApp(const IslandChatApp());
}

class IslandChatApp extends StatelessWidget {
  const IslandChatApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'IslandChat',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.green),
        useMaterial3: true,
      ),
      home: const ChatScreen(),
      debugShowCheckedModeBanner: false, // Treu la pestanya de "DEBUG"
    );
  }
}